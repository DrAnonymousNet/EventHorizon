from djoser.serializers import  UserCreatePasswordRetypeSerializer, UserCreateSerializer as BaseUserCreateSerializer
from rest_framework import serializers
from apps.account.models import Device
from apps.core.api.serializers import RelationSerializerMixin
from django.contrib.auth import get_user_model
from djoser.serializers import UserCreatePasswordRetypeSerializer


User = get_user_model()

class DeviceSerializer(RelationSerializerMixin, serializers.ModelSerializer):
    class Meta:
        model = Device
        fields = [
            "device_token"
        ]
        relation_fields = {
            "user": User
        }

class UserCreateSerializer(BaseUserCreateSerializer):
    devices = DeviceSerializer(many=True, required=False)

    class Meta(BaseUserCreateSerializer.Meta):
        model = get_user_model()
        fields = BaseUserCreateSerializer.Meta.fields + ("uid","devices")#, "username")
        BaseUserCreateSerializer.Meta.extra_kwargs.update({
            "username":{
                "required":True
            }
        })

    def create(self, validated_data):
        return super().create(validated_data)


class UserCreatePasswordRetypeSerializer(UserCreatePasswordRetypeSerializer, UserCreateSerializer):
    pass         
