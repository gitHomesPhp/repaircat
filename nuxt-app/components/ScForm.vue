<template>
  <form name="addScForm" @submit.prevent.stop="addSc" class="sc-form">
    <label>
      <span class="title">Название:</span>
      <input type="text" size="50" v-model="sc.name">
    </label>
    <label>
      <span class="title">Описание:</span>
      <textarea v-model="sc.description" cols="49" rows="10"></textarea>
    </label>
    <label>
      <span>Телефон:</span>
      <input type="text" size="50" v-model="sc.phone">
    </label>
    <label>
      <span>Почта:</span>
      <input type="text" size="50" v-model="sc.email">
    </label>
    <label>
      <span>Сайт:</span>
      <input type="text" size="50" v-model="sc.site">
    </label>
    <label>
      <input type="submit" value="Добавить">
    </label>
  </form>
</template>

<script setup lang="ts">
  const config = useRuntimeConfig()
  const sc = ref({
    name: '',
    description: '',
    phone: '',
    email: '',
    site: '',
  })

  const addSc = async () => {
    const data = new FormData()
    data.append('name', sc.value.name)
    data.append('description', sc.value.description)
    data.append('phone', sc.value.phone)
    data.append('email', sc.value.email)
    data.append('site', sc.value.site)

    await $fetch(config.API + '/sc', {
      body: data,
      method: 'post'
    })
      .then(() => {
        sc.value.name = ''
        sc.value.description = ''
        sc.value.phone = ''
        sc.value.email = ''
        sc.value.site = ''
      })
      .catch((err) => console.log(err))
  }
</script>

<style scoped lang="scss">
  textarea {
    resize: none;
  }
  .sc-form {
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