# -*- coding: utf-8 -*-

from locust import HttpUser, task, between

class WebsiteUser(HttpUser):
    host = 'http://api-svc.api.svc.cluster.local:8888'
    wait_time = between(1, 2)

    @task
    def get_hello_world(self):
        self.client.get('/api/v1/sayHello')