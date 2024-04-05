from rest_framework import routers

from apps.event.api.v1.views import EventAPIViewSet

router = routers.SimpleRouter()
router.register(r"event", EventAPIViewSet)

urlpatterns = [] + router.urls
