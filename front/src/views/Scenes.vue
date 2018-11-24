<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>VROMStart</th>
          <th>VROMEnd</th>
          <th>ChestPositionsCount</th>
        </tr>
      </thead>

      <tbody>
          <tr v-for="scene in scenes">
            <td>{{ scene.Name }}</td>
            <td>{{ scene.VROMStart | hex(8) }}</td>
            <td>{{ scene.VROMEnd | hex(8) }}</td>
            <td>{{ scene.ChestPositionsCount }}</td>
          </tr>
        </v-for>
      </tbody>

    </table>
  </div>
</template>

<script src="./Scenes.js"></script>

<style scoped>
table {
  width: 100%;
}
</style>
