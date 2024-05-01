import json

from apps.account.models import Device
from apps.core.api.serializers import RelationSerializerMixin
from config.publisher import publisher
from django.contrib.auth import get_user_model
from djoser.serializers import (
    UserCreatePasswordRetypeSerializer as BaseUserCreatePasswordRetypeSerializer,
)
from djoser.serializers import UserCreateSerializer as BaseUserCreateSerializer
from rest_framework import serializers

User = get_user_model()


class DeviceSerializer(RelationSerializerMixin, serializers.ModelSerializer):
    class Meta:
        model = Device
        fields = ["device_token"]
        relation_fields = {"user": User}

    def create(self, validated_data):
        return super().create(validated_data)


class UserCreateSerializer(BaseUserCreateSerializer):
    devices = DeviceSerializer(many=True, required=False)

    class Meta(BaseUserCreateSerializer.Meta):
        model = get_user_model()
        fields = BaseUserCreateSerializer.Meta.fields + ("uid", "devices", "username")
        extra_kwargs = {"username": {"required": True}}

    def validate(self, attrs):
        # Djoser calls User(**attrs) which result in
        # TypeError: Direct assignment to the reverse side of a related
        # set is prohibited. Use devices.set() instead.TypeError: Direct assignment to the reverse side of a related set is prohibited. Use devices.set() instead.TypeError: Direct assignment to the reverse side of a related set is prohibited. Use devices.set() instead.TypeError: Direct assignment to the reverse side of a related set is prohibited. Use devices.set() instead.
        devices = attrs.pop("devices")
        atrrs = super().validate(attrs)
        atrrs["devices"] = devices
        return atrrs

    def create(self, validated_data):
        devices = validated_data.pop("devices")
        user = super().create(validated_data)
        for device in devices:
            device.update({"user": user})
            Device.objects.create(**device)
        publisher.publish_message(json.dumps(self.data))
        return user


class UserCreatePasswordRetypeSerializer(
    BaseUserCreatePasswordRetypeSerializer, UserCreateSerializer
):
    pass
