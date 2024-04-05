from rest_framework.serializers import ModelSerializer
from rest_framework import serializers

from apps.core.api.serializers import RelationSerializerMixin
from apps.event.models import Category, Event, EventAttendance, EventImage, Tag


class CategorySerializer(ModelSerializer):
    class Meta:
        model = Category
        fields = "__all__"


class TagSerializer(ModelSerializer):
    class Meta:
        model = Tag
        fields = "__all__"


class EventImageSerializer(RelationSerializerMixin, ModelSerializer):
    class Meta:
        model = EventImage
        fields = (
            "uid",
            "created",
            "updated",
            "event",
            "caption",
            "image",
        )
        relation_fields = {"event": Event}
        relation_kwargs = {"event": {"read_only": True}}


class EventAttendanceSerializer(RelationSerializerMixin, ModelSerializer):
    class Meta:
        model = EventAttendance
        fields = [
            "uid",
            "created",
            "updated",
            "event",
            "unregistered_user_email",
            "registered_user",
            "attendance_type",
        ]
        relation_fields = {"event": Event}
        relation_kwargs = {"event": {"read_only": True}}



class EventFullSerializer(ModelSerializer):
    images = EventImageSerializer(many=True)
    categories = CategorySerializer(many=True)
    attendees = EventAttendanceSerializer(many=True, read_only=True)

    class Meta:
        fields = [
            "uid",
            "title",
            "description",
            "organizer",
            "attendees",
            "location",
            "start_time",
            "end_time",
            "categories",
            "tags",
            "max_participants",
            "is_virtual",
            "virtual_meeting_link",
            "status",
        ]
        extra_kwargs = {"organizer": {"source": "email"}}


class EventMinimalSerializer(ModelSerializer):
    class Meta:
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
