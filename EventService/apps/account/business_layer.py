from config.publisher import rbmq_publisher


def publish_user_created_message(user_created_response):
    rbmq_publisher(user_created_response)
