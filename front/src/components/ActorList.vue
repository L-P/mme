<template>
  <table class="table">

    <thead>
      <th>ID</th>
      <th>Name</th>
      <th>Description</th>
      <th>Initialization</th>
      <th>SpawnTimeFlags</th>
    </thead>

    <tbody>
      <tr v-for="v, k in actors" :key="k">
        <td>{{v.ID | hex(4)}}</td>
        <td>{{v.Description.FileName}}</td>
        <td>{{[v.Description.Identification, v.Description.Translation] | coalesce}}</td>
        <td>{{v.Initialization | hex(4)}}</td>
        <td>{{v.SpawnTimeFlags | hex(4)}}</td>
      </tr>
    </tbody>

  </table>
</template>

<script src="./ActorList.js"></script>

<style scoped>
.table {
  width: 100%;
}
</style>
