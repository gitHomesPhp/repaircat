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
      <select v-model="sc.city">
        <option v-for="city in cityMenu" :value="city">{{ city.label }}</option>
      </select>
    </label>
    <label>
      <span>Адресс:</span>
      <input type="text" size="50" v-model="sc.address">
    </label>
    <!--label>
      <span>Метро:</span>
      <select v-model="sc.underground">
        <option v-for="underground in undergrounds" :value="underground">{{ underground.label }}</option>
      </select>
    </label-->
    <label>
      <span>Метро:</span>
      <span class="undergrounds">
        <span v-for="(subway, i) in sc.undergrounds">
          <select v-model="sc.undergrounds[i]" :key="i">
            <option v-for="underground in undergrounds" :value="underground">{{ underground.label }}</option>
          </select>
          <button @click.prevent="removeUnderground(i)" class="pointer">-</button>
        </span>
        <button @click.prevent="addUnderground" class="pointer">+</button>
      </span>
    </label>
    <label>
      <span>Широта:</span>
      <input type="text" size="50" v-model="sc.latitude">
    </label>
    <label>
      <span>Долгота:</span>
      <input type="text" size="50" v-model="sc.longitude">
    </label>
    <div>
      <label>
        <input type="submit" value="Добавить">
      </label>
    </div>
  </form>
</template>

<script setup lang="ts">
  const route = useRoute()
  const undergrounds = ref([])
  const NO_UNDERGROUND = 0

  const cityMenu = await $fetch('/api/city')

  const sc = ref({
    name: '',
    description: '',
    phone: '',
    email: '',
    site: '',
    city: cityMenu[0],
    address: '',
    //underground: {},
    undergrounds: [],
    latitude: '',
    longitude: '',
  })

  undergrounds.value = await $fetch(`/api/underground?city=${cityMenu[0].id}`)

  const addSc = async () => {
    await $fetch('/api/sc', {
      body: {
        name: sc.value.name,
        description: sc.value.description,
        phone: sc.value.phone,
        email: sc.value.email,
        site: sc.value.site,
        city: sc.value.city.id,
        address: sc.value.address,
        //underground: sc.value.underground.id ?? NO_UNDERGROUND,
        latitude: sc.value.latitude,
        longitude: sc.value.longitude,
        undergrounds: sc.value.undergrounds,
        user: route.query.user,
        token: route.query.token,
      },
      method: 'post'
    })
      .then(() => {
        //sc.value.name = ''
        //sc.value.description = ''
        //sc.value.phone = ''
        //sc.value.email = ''
        //sc.value.site = ''
        //sc.value.address = ''
        //sc.value.underground = ''
      })
      .catch((err) => console.log(err))
  }

  const addUnderground = () => {
    sc.value.undergrounds.push({})
  }

  const removeUnderground = (i) => {
    sc.value.undergrounds.splice(i, 1);
  }

  watch(() => sc.value.city, async (value) => {
    undergrounds.value = await $fetch(`/api/underground?city=${value.id}`)
  })

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
      input {
        width: 30rem;
      }
      textarea {
        width: 30rem;
      }
      select {
        width: 30.5rem;
      }
    }
  }
  @media (max-width: 670px) {
    input {
      width: 23rem;
    }
    textarea {
      width: 23rem;
    }
    select {
      width: 23.5rem;
    }
  }
  .undergrounds {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    select {
      margin-bottom: .2rem;
    }
  }
</style>