from rest_framework.serializers import ModelSerializer

from apps.event.models import Category, Event, EventAttendance, EventImage, Tag


class CategorySerializer(ModelSerializer):
    class Meta:
        model = Category
        fields = "__all__"


class TagSerializer(ModelSerializer):
    class Meta:
        model = Tag
        fields = "__all__"


class EventImageSerializer(ModelSerializer):
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


class EventAttendanceSerializer(ModelSerializer):
    class Meta:
        model = EventAttendance
        fields = "__all__"
