# Golang GraphQL MongoDB CRUD Project

This project provides a complete Golang implementation for a CRUD (Create, Read, Update, Delete) API using GraphQL and MongoDB. It leverages the powerful `gqlgen` library for generating a type-safe GraphQL schema and handlers.

**Features:**

* **GraphQL API:** Exposes a clean and well-defined API for managing Jobs.
* **MongoDB Integration:** interacts with MongoDB for data storage and retrieval.
* **CRUD Operations:** Supports creation, retrieval, update, and deletion of Jobs.

**Getting Started:**

1. **Prerequisites:**
  - Go and Git installed on your system.
2. **Project Setup:**
  - Clone this repository: `git clone https://github.com/jay-SP/gql.git`
  - Install dependencies: `go mod tidy`
3. Get gql gen for your project go get github.com/99designs/gqlgen

**Interaction with the API:**
The project utilizes a GraphQL schema to define the available queries and mutations. These queries and mutations allow you to interact with your Job data through your preferred GraphQL client.

**Example Queries:**

- **Get All Jobs:**
  ```graphql
  query GetAllJobs {
    jobs {
      _id
      title
      description
      company
      url
    }
  }

- **Create Job**
  ```graphql
  mutation CreateJobListing($input: CreateJobListingInput!) {
  createJobListing(input: $input) {
    _id
    title
    description
    company
    url
  }
  }
  {
  "input": {
    "title": "Software Development Engineer - I",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do   
     eiusmod tempor incididunt",
    "company": "Google",
    "url": "[www.google.com/](https://www.google.com/)"  
     }
   }

- **Get Job by ID**
  ```graphql
    query GetJob($id: ID!) {
    job(id: $id) {
      _id
      title
      description
      company
      url
      }
    }
    {
      "id": "638051d7acc418c13197fdf7"
    }

- **Update Job**
  ```graphql
  mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) {
    updateJobListing(id: $id, input: $input) {
      title
      description
      _id
      company
      url
      }
    }
    {
    "id": "638051d3acc418c13197fdf6",
    "input": {
    "title": "Software Development Engineer - III"
      }
    }  

- **Delete Job by ID**
```graphql
mutation DeleteQuery($id: ID!) {
  deleteJobListing(id: $id) {
    deletedJobId
  }
}

{
  "id": "638051d3acc418c13197fdf6"
}
