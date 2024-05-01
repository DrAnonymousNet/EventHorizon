from apps.core.api.views import ResponseViewMixin
from apps.core.utils import str_to_bool
from apps.event.api.v1.serializers import (
    EventAttendanceSerializer,
    EventFullSerializer,
    EventMinimalSerializer,
    EventRSVPSerializer,
)
from apps.event.business_layer import send_rsvp_message_to_notification_service
from apps.event.models import Event
from drf_yasg import openapi
from drf_yasg.utils import swagger_auto_schema
from rest_framework import status
from rest_framework.decorators import action
from rest_framework.parsers import JSONParser, MultiPartParser
from rest_framework.permissions import SAFE_METHODS, AllowAny, IsAuthenticatedOrReadOnly
from rest_framework.response import Response
from rest_framework.viewsets import ModelViewSet


class EventAPIViewSet(ResponseViewMixin, ModelViewSet):
    serializer_class = EventFullSerializer
    queryset = Event.objects.all()
    permission_classes = [
        IsAuthenticatedOrReadOnly,
    ]
    serializer_class = EventFullSerializer
    parser_classes = [JSONParser, MultiPartParser]
    lookup_field = "uid"

    def get_serializer_class(self):
        full_response = str_to_bool(self.request.query_params.get("full", "false"))
        if full_response or self.request.method not in SAFE_METHODS:
            return super().get_serializer_class()
        return EventMinimalSerializer

    @swagger_auto_schema(
        manual_parameters=[
            openapi.Parameter(
                "full",
                in_=openapi.IN_QUERY,
                description="Return full response",
                type=openapi.TYPE_BOOLEAN,
            )
        ]
    )
    def retrieve(self, request, *args, **kwargs):
        return super().retrieve(request, *args, **kwargs)

    @swagger_auto_schema(
        manual_parameters=[
            openapi.Parameter(
                "full",
                in_=openapi.IN_QUERY,
                description="Return full response",
                type=openapi.TYPE_BOOLEAN,
            )
        ]
    )
    def list(self, request, *args, **kwargs):
        return super().list(request, *args, **kwargs)

    @swagger_auto_schema(
        request_body=EventRSVPSerializer(), responses={200: EventAttendanceSerializer()}
    )
    @action(
        methods=["post"],
        url_path="rsvp",
        detail=True,
        authentication_classes=[],
        permission_classes=[AllowAny],
    )
    def rsvp(self, request, *args, **kwargs):
        event = self.get_object()
        serializer = EventRSVPSerializer(
            data=request.data, context={"event": event, "request": request}
        )
        if not serializer.is_valid():
            return Response(
                status=status.HTTP_400_BAD_REQUEST,
                data={"msg": "validation failed", **serializer.errors},
            )
        event_attendee = serializer.rsvp()
        send_rsvp_message_to_notification_service(event, event_attendee)
        return Response(
            status=status.HTTP_200_OK,
            data={"msg": "You have rsvp successfuly", **serializer.data},
        )
