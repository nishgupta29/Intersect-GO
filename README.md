# Intersect-GO
 Go learning repository
This repository contains my work towards learning GO lang and also the work towards Intersect assignment.

First few commits are basically for learnning GO where i created a simple hello world code and then a simple rest API that displays "Welcome to Intersect" in browser on port 8000.

Then i tried creating a simple slack bot in GO, which connects to slack using a API token which is stored in an environment variable "SLACK_API"
This bot reads a reply from user "hello name" and then displays "hello" in return.

I e dockerised this application using a simple dockerfile which uses golang as a base image.
