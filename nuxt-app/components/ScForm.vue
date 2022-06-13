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
      <span>Город:</span>
      <input type="text" size="50" v-model="sc.city">
    </label>
    <label>
      <span>Адресс:</span>
      <input type="text" size="50" v-model="sc.address">
    </label>
    <label>
      <span>Метро:</span>
      <input type="text" size="50" v-model="sc.underground">
    </label>
    <label>
      <input type="submit" value="Добавить">
    </label>
  </form>
</template>

<script setup lang="ts">
  const config = useRuntimeConfig()
  const route = useRoute()

  const sc = ref({
    name: '',
    description: '',
    phone: '',
    email: '',
    site: '',
    city: '',
    address: '',
    underground: '',
  })

  const addSc = async () => {
    const data = new FormData()
    data.append('name', sc.value.name)
    data.append('description', sc.value.description)
    data.append('phone', sc.value.phone)
    data.append('email', sc.value.email)
    data.append('site', sc.value.site)
    data.append('city', sc.value.city)
    data.append('address', sc.value.address)
    data.append('underground', sc.value.underground)
    data.append('user', route.query.user)
    data.append('token', route.query.token)

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
        sc.value.city = ''
        sc.value.address = ''
        sc.value.underground = ''
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