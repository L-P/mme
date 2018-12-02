<template>
  <table class="table">

    <thead>
      <th>ID</th>
      <th>Initialization</th>
      <th>SpawnTimeFlags</th>
      <th>RotationFlags</th>
      <th>Position</th>
      <th>Rotation</th>
    </thead>

    <tbody>
      <tr v-for="v, k in actors" :key="k">
        <td>{{v.ID | hex(4)}}</td>
        <td>{{v.Initialization | hex(4)}}</td>
        <td>{{v.SpawnTimeFlags | hex(4)}}</td>
        <td>
            {{v.NoXRotation | bool}}
            {{v.NoYRotation | bool}}
            {{v.NoZRotation | bool}}
        </td>
        <td>{{v.Position}}</td>
        <td>{{v.Rotation}}</td>
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
