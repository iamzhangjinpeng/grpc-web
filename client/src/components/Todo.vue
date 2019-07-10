<template>
  <div class="container todo">
    <h1>Todo List</h1>

    <table class="table table-bordered">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">Task</th>
          <th scope="col">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(todo, index) in todos" :key="index">
          <td>{{ todo.id }}</td>
          <td>{{ todo.task }}</td>
          <td>
            <button type="button" class="btn btn-light btn-sm" @click="deleteTodo(index)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>

    <form>
      <div class="form-group">
        <label for="Task">New Task</label>
        <input v-model.trim="task" type="text" class="form-control" id="Task" placeholder="Enter Task">
      </div>
      <button type="button" class="btn btn-primary btn-sm" @click="addTodo">Submit</button>
    </form>
  </div>
</template>

<script>
import "bootstrap/dist/css/bootstrap.css"
import { listTodoParams, addTodoParams, deleteTodoParams } from "@/lib/todo_pb"
import { todoServiceClient } from "@/lib/todo_grpc_web_pb"

const client = new todoServiceClient("http://localhost:8080", null, null)

export default {
  name: 'Todo',
  data: () => {
    return {
      todos: [],
      task: ''
    }
  },
  created () {
    this.listTodo()
  },
  methods: {
    reset () {
      this.task = ''
    },
    listTodo () {
      let params = new listTodoParams()
      client.listTodo(params, {}, (err, response) => {
        console.log('listTodo', response.toObject())
        this.todos = response.toObject().todosList
      })
    },
    addTodo () {
      if (this.task) {
        let params = new addTodoParams()
        params.setTask(this.task)
        client.addTodo(params, {}, (err, response) => {
          console.log('addTodo', response.toObject())
          this.listTodo()
          this.reset()
        })
      }
    },
    deleteTodo (index) {
      let id = this.todos[index].id
      let params = new deleteTodoParams()
      params.setId(id)
      client.deleteTodo(params, {}, (err, response) => {
        console.log('deleteTodo', response.toObject())
        this.listTodo()
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.todo {
  margin-top: 50px;
}
</style>
