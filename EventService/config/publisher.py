import json
import time

import pika
from django.conf import settings
from pika.exceptions import AMQPConnectionError


class RabbitMQPublisher:
    def __init__(self, host="localhost"):
        self.connection = None
        self.channel = None
        self.host = host
        self.connect()

    def connect(self):
        try:
            self.connection = pika.BlockingConnection(
                pika.ConnectionParameters(self.host)
            )
            self.channel = self.connection.channel()
            self.channel.exchange_declare(
                exchange="userCreated",
                exchange_type="fanout",
                durable=True,
                auto_delete=False,
                internal=False,
            )
        except AMQPConnectionError:
            print("Connection failed, retrying...")
            # TODO Add exponential backoff ?
            time.sleep(5)
            self.connect()  # Recursive reconnection attempt

    def publish_message(self, message):
        if not self.connection or not self.connection.is_open:
            print("Connection closed, reconnecting...")
            self.connect()
        try:
            exchange = "userCreated"
            routing_key = "userCreated"
            self.channel.basic_publish(
                exchange=exchange, routing_key=routing_key, body=json.dumps(message)
            )
            print("Message published")
        except (AMQPConnectionError, Exception) as e:
            print(f"Publishing failed: {e}")
            # Handle specific exceptions and reconnection logic here


# Usage
publisher = RabbitMQPublisher()
