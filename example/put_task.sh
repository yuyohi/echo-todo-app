#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"sample task", "description":"サンプルのタスクです(更新)", "status": "todo"}' localhost:1234/tasks/1