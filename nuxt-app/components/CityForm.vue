<template>
  <form @submit.prevent.stop="addCity" class="city-form">
    <label>
      <span class="title">Название:</span>
      <input type="text" v-model="city.label">
    </label>
    <label>
      <span class="title">Код:</span>
      <input type="text" v-model="city.code">
    </label>
    <label>
      <input type="submit" value="Добавить">
    </label>
  </form>
</template>

<script setup lang="ts">
  const route = useRoute()

  const city = ref({
    label: '',
    code: '',
  })

  const addCity = async () => {
    await $fetch('api/city', {
      body: {
        label: city.value.label,
        code: city.value.code,
        user: route.query.user,
        token: route.query.token,
      },
      method: 'post'
    })
  }
</script>

<style scoped lang="scss">
.city-form {
  display: flex;
  flex-direction: column;
  label {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: .5rem;
  }
}
</style>