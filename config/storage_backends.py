from django.conf import settings
from django.core.files.storage import FileSystemStorage
from django.utils.deconstruct import deconstructible
from storages.backends.s3boto3 import S3Boto3Storage

if not settings.DEBUG:

    @deconstructible
    class StaticStorage(S3Boto3Storage):
        location = "static"
        default_acl = "public-read"
        bucket_name = settings.AWS_STORAGE_PUBLIC_BUCKET_NAME
        custom_domain = f"{settings.AWS_STORAGE_PUBLIC_BUCKET_NAME}.s3.amazonaws.com"

    @deconstructible
    class PublicMediaStorage(S3Boto3Storage):
        location = "media"
        default_acl = "public-read"
        file_overwrite = False
        bucket_name = settings.AWS_STORAGE_PUBLIC_BUCKET_NAME
        custom_domain = f"{settings.AWS_STORAGE_PUBLIC_BUCKET_NAME}.s3.{settings.AWS_STORAGE_PUBLIC_REGION_NAME}.amazonaws.com"

    @deconstructible
    class PrivateMediaStorage(S3Boto3Storage):
        location = "private"
        default_acl = "private"
        file_overwrite = False
        custom_domain = False

else:

    @deconstructible
    class StaticStorage(FileSystemStorage):
        pass

    @deconstructible
    class PublicMediaStorage(FileSystemStorage):
        pass

    @deconstructible
    class PrivateMediaStorage(FileSystemStorage):
        pass
