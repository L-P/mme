<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Offset</th>
          <th>VROMStart</th>
          <th>String</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="message in messages"
          :key="message.ID"
        >
          <td>{{ message.ID | hex(4) }}</td>
          <td>{{ message.Offset | hex(8) }}</td>
          <td>{{ message.VROMStart | hex(8) }}</td>
          <td style="white-space: pre;">
            {{ message.String }}
          </td>
        </tr>
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
