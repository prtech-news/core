#!/bin/bash

aws lambda invoke --function-name spider --payload file://payload.json --cli-binary-format raw-in-base64-out out.json