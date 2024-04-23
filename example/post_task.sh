#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"sample task", "description":"サンプルのタスクです", "status": "todo"}' localhost:1234/tasks
