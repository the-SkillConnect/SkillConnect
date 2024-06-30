# SkillConnect
A Freelancing Telegram Mini App (TMA)

# How to test out the GraphQL API
## Insert a user
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation InsertUser { insertUser(input: { firstname: \"milad\", surname: \"rrr\", mobilePhone: \"5168416\", email: \"mili@gmail.com\", password: \"1234\" }) { email firstname id mobilePhone password surname } }"
  }' \
  http://localhost:8585/graphql
```

## Insert a project
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation InsertProject { insertProject(input: { total_amount: \"123.1\", status: false, user_id: 1, fee: \"84.01\", title: \"tit\", description: \"dec\" }) { description fee id orderDate status title total_amount user_id } }"
  }' \
  http://localhost:8585/graphql

```

## Get a project

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query Projects { projects { description fee id orderDate status title total_amount user_id } }"
  }' \
  http://localhost:8585/graphql

```

## Update the Project

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation UpdateProject { updateProject(input: { title: \"mili\", description: \"descrip2\", total_amount: \"12\", status: true, user_id: 1, fee: \"13\", id: 1 }) { description fee id orderDate status title total_amount user_id } }"
  }' \
  http://localhost:8585/graphql

```

## Insert a project comment

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation InsertProjectComment { insertProjectComment(input: { user_id: 1, project_id: 1, text: \"1635115\" }) { date id project_id text user_id } }"
  }' \
  http://localhost:8585/graphql

```

## Update a project comment:

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation UpdateProjectComment { updateProjectComment(input: { text: \"bad text is here\", id: 5, user_id: 4, project_id: 2 }) { email firstname id mobilePhone password surname } }"
  }' \
  http://localhost:8585/graphql


```