<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Number</th>
          <th>Offset</th>
          <th>VROMStart</th>
          <th>String</th>
        </tr>
      </thead>

      <tbody>
          <tr v-for="message in messages">
            <td>{{ message.Number | hex(4) }}</td>
            <td>{{ message.Offset | hex(8) }}</td>
            <td>{{ message.VROMStart | hex(8) }}</td>
            <td style="white-space: pre;">{{ message.String }}</td>
          </tr>
        </v-for>
      </tbody>

    </table>
  </div>
</template>

<script src="./Messages.js"></script>

<style scoped>
table {
  width: 100%;
}
</style>
