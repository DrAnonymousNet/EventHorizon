from django.conf import settings
from django.conf.urls import include
from django.urls import path, re_path
from rest_framework.routers import DefaultRouter, SimpleRouter

if settings.DEBUG:
    router = DefaultRouter()
else:
    router = SimpleRouter()

app_name = "api"
urlpatterns = router.urls

urlpatterns += [
    path("", include("apps.event.api.urls")),
    re_path(r'^auth/', include('djoser.urls')),

]