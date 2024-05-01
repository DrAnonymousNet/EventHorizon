from apps.core.api.serializers import RelationSerializerMixin
from apps.event.models import Event, EventAttendance, EventImage
from django.contrib.auth import get_user_model
from django.db.models import Q
from drf_writable_nested import WritableNestedModelSerializer
from rest_framework import serializers

User = get_user_model()


class EventImageSerializer(RelationSerializerMixin, serializers.ModelSerializer):
    class Meta:
        model = EventImage
        fields = (
            "uid",
            "created",
            "modified",
            "event",
            "caption",
            "image",
        )
        relation_fields = {"event": Event}
        relation_kwargs = {"event": {"read_only": True}}


class EventAttendanceSerializer(RelationSerializerMixin, serializers.ModelSerializer):
    registered_user = serializers.SlugRelatedField(
        slug_field="email", queryset=User.objects.all()
    )

    class Meta:
        model = EventAttendance
        fields = [
            "uid",
            "name",
            "created",
            "modified",
            "event",
            "unregistered_user_email",
            "registered_user",
            "attendance_type",
        ]
        relation_fields = {"event": Event}
        relation_kwargs = {"event": {"read_only": True}}


class EventRSVPSerializer(serializers.Serializer):
    AttendanceType = EventAttendance.AttendanceType
    email = serializers.EmailField()
    attendance_type = serializers.ListField(
        child=serializers.ChoiceField(choices=AttendanceType.choices)
    )

    def validate(self, attrs):
        request = self.context.get("request")
        user = request.user
        if isinstance(user, User):
            attrs["email"] = user.email
        elif not attrs.get("email", None):
            raise serializers.ValidationError("Email is required to RSVP or Login")

        event = self.context.get("event")
        email = attrs.get("email")
        attendee_already_rsvp = event.attendees.filter(
            Q(unregistered_user_email=email) | Q(registered_user__email=email)
        )
        if attendee_already_rsvp:
            raise serializers.ValidationError("you already RSVP for this event")

        return super().validate(attrs)

    def rsvp(self):
        event = self.context.get("event")
        event_attendance = EventAttendance(
            event=event, attendance_type=self.validated_data.get("attendance_type")
        )
        email = self.validated_data.get("email")
        try:
            event_attendance.registered_user = User.objects.get(email=email)
        except User.DoesNotExist:
            event_attendance.unregistered_user_email = email
        event_attendance.save()
        return event_attendance


class EventFullSerializer(RelationSerializerMixin, WritableNestedModelSerializer):
    images = EventImageSerializer(many=True)
    attendees = EventAttendanceSerializer(many=True, read_only=True)

    class Meta:
        model = Event
        fields = [
            "uid",
            "title",
            "description",
            "organizer",
            "attendees",
            "location",
            "start_time",
            "end_time",
            "max_participants",
            "is_virtual",
            "virtual_meeting_link",
            "status",
            "images",
        ]
        relation_fields = {"organizer": User}

    def validate(self, attrs):
        request = self.context.get("request")
        attrs["organizer"] = request.user
        return super().validate(attrs)


class EventMinimalSerializer(serializers.ModelSerializer):
    class Meta:
        model = Event
        fields = [
            "title",
            "description",
            "organizer",
            "location",
            "start_time",
            "end_time",
            "max_participants",
            "is_virtual",
            "virtual_meeting_link",
            "status",
        ]

        relation_fields = {"organizer": User}
