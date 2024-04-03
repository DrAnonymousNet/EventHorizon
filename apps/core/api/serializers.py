import collections

from django.utils.translation import gettext_lazy as _
from rest_framework import serializers
from rest_framework.exceptions import NotFound, ValidationError


class ForceQuerySerializer(serializers.Serializer):
    force = serializers.BooleanField(default=False)


class NoModelSerializer(serializers.Serializer):
    """
    Normal serializer, not linked to a model
    This serializer should not call the `save()` method
    So it overrides `create()` and `update()` to raise `RuntimeError` when executed
    """

    def create(self, validated_data):
        raise RuntimeError("`create()` should not be called.")

    def update(self, instance, validated_data):
        raise RuntimeError("`update()` should not be called.")


class RelationSerializerMixin:
    """
    By default, relation fields are represented using the model id
    This mixin represents relation fields using the model uid
    It also accepts the uid in request body and converts it to the model instance

    `usage`
    class CustomSerializer(RelationSerializerMixin, serializers.Serializer):
        class Meta:
            relation_fields = {"field": Model}
            relation_kwargs = {"field": {"read_only": True}}  # optional
    """

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        name = self.__class__.__name__
        assert hasattr(self, "Meta"), f"Missing `Meta` in `{name}`"
        assert hasattr(
            self.Meta, "relation_fields"
        ), f"Missing `relation_fields` in `{name}.Meta`"

        mapping = self.Meta.relation_fields
        assert isinstance(
            mapping, collections.Mapping
        ), f"`{name}.Meta.relation_fields` must be of type `Mapping`"

        fields = mapping.keys()

        extra_kwargs = {}
        if hasattr(self.Meta, "relation_kwargs"):
            extra_kwargs = self.Meta.relation_kwargs
            assert isinstance(
                extra_kwargs, collections.Mapping
            ), f"`{name}.Meta.relation_kwargs` must be of type `Mapping`"

        for field in fields:
            self.fields[field] = serializers.UUIDField(
                source=f"{field}.uid",
                default=None,
                **extra_kwargs.get(field, {}),
            )

    def validate(self, data):
        for field, model in self.Meta.relation_fields.items():
            if field in data:
                # Already fetched from DB, e.g. (Nested serializer)
                if isinstance(data[field], model):
                    continue

                # Relation "uid" was sent by client
                if "uid" in data[field]:
                    # Set relation when "uid" has value
                    if data[field]["uid"]:
                        try:
                            data[field] = model.objects.get(uid=data[field]["uid"])
                        except model.DoesNotExist:
                            raise NotFound(f"{model.__name__} does not exist")
                    # Clear relation when "uid" is None
                    else:
                        data[field] = None

                # Relation "uid" was not sent by client, e.g. (PATCH request)
                else:
                    data.pop(field)
        return super().validate(data)


class UIDListSerializerMixin:
    """
    Adds `uid_list` field to the serializer

    class CustomSerializer(UIDListSerializerMixin, serializers.Serializer):
        class Meta:
            uid_list_kwargs = {"help_text": "Descriptive text"}
    """

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        name = self.__class__.__name__
        assert hasattr(self, "Meta"), f"Missing `Meta` in `{name}`"
        assert hasattr(
            self.Meta, "uid_list_kwargs"
        ), f"Missing `uid_list_kwargs` in `{name}.Meta`"

        extra_kwargs = self.Meta.uid_list_kwargs
        assert isinstance(
            extra_kwargs, collections.Mapping
        ), f"`{name}.Meta.uid_list_kwargs` must be of type `Mapping`"
        assert (
            "help_text" in extra_kwargs.keys()
        ), f"Missing `help_text` in `{name}.Meta.uid_list_kwargs`"

        self.fields["uid_list"] = serializers.ListSerializer(
            child=serializers.UUIDField(), **extra_kwargs
        )


class UIDListSerializer(UIDListSerializerMixin, NoModelSerializer):
    class Meta:
        uid_list_kwargs = {"help_text": "UID List"}
