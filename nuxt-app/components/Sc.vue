<template>
  <div class="sc">
    <div class="sc__header">
      <div class="sc__title">
        <div class="sc__type">
          <span>{{ title }}</span>
        </div>
        <h1 class="sc__name">{{sc.name}}</h1>
      </div>
      <div class="sc__contacts contacts">
        <div class="contact__column contact__column--header">
          <span>Телефон:</span>
          <span>Почта:</span>
          <span>Сайт:</span>
        </div>
        <div class="contact__column contact__column--value">
          <span class="contact__item">{{sc.phone || 'Нет информации'}}</span>
          <span class="contact__item">{{sc.email || 'Нет информации'}}</span>
          <span class="contact__item">{{sc.site || 'Нет информации'}}</span>
        </div>
      </div>
    </div>
    <div class="sc__body">
      <div class="sc__meta-info">
        <div>рейтинг 1000</div>
      </div>
      <div class="sc__line"></div>
      <div class="sc__main-info">
        <div @click="toggleText" class="sc__description pointer" :class="{'hidden': isHidden}">
          <div  v-html="sc.description"></div>
        </div>
      </div>
    </div>
    <div class="sc__address">
      <div>
        <div class="sc__location-address location-info">
          <img src="~/assets/img/location_icon.png" alt="">
          <span class="location-info__header">Аддрес:</span>
          <span class="location-info__value">{{ sc.location.address }}</span>
        </div>
        <div v-if="sc.location.undergrounds.length"
             v-for="underground in sc.location.undergrounds"
             class="sc__location-address location-info"
        >
          <img src="~/assets/img/underground_icon.png" alt="">
          <span class="location-info__header">Метро:</span>
          <span class="location-info__value">{{ underground.label }}</span>
        </div>
      </div>
      <div class="sc__map">
        <client-only>
          <yandex-map
              v-if="showMap"
              :coords="coords"
              zoom=16
              @click="changeCoords"
          >
            <ymap-marker
                :coords="coords"
                marker-id="123123"
                hint-content="some hint"
            />
          </yandex-map>
        </client-only>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  const route = useRoute()

  const title = ref('Сервисный центр')
  const { data: sc} = await useFetch(`/api/sc/${route.params.id}`)
  const isHidden = ref(true)
  const toggleText = () => {isHidden.value = !isHidden.value}
  const showMap  = ref(false)

  const coords = ref([sc.value.location.latitude, sc.value.location.longitude])
  const changeCoords = (e) => {
    coords.value =  e.get('coords');
  }
  onMounted(() => {
    setTimeout(() => {showMap.value = true}, 150)

  })
</script>

<style scoped lang="scss">
  .sc {
    display: flex;
    flex-direction: column;
    &__header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding-bottom: .5rem;
      @media (max-width: 510px) {
        flex-direction: column;
      }
      img {
        width: 200px;
      }
    }
    &__contacts {
      display: flex;

    }
    &__type{
      padding: .3rem 0;
      color: #6c757d;
    }
    &__name {
      font-size: 2rem;
      font-weight: bold;
    }
    &__title {
      @media (max-width: 510px) {
        padding-bottom: 1rem;
      }
    }
    &__body {
      display: flex;
      border-top: #cccccc 1px solid;
      border-bottom: #cccccc 1px solid;
      justify-content: space-around;
      @media (max-width: 510px) {
        flex-direction: column-reverse;
      }
    }
    &__main-info {
      display: flex;
      flex-direction: column;
      max-width: 66%;
      color: #5a5c61;
      font-family: Roboto,Arial,sans-serif;
      font-size: 1rem;;
      padding-top: .5rem;
      padding-bottom: .5rem;
      @media (max-width: 510px) {
        max-width: 100%;
      }
    }
    &__line {
      border-right: #cccccc 1px solid;
      @media (max-width: 510px) {
        border-right: none;
        border-bottom: #cccccc 1px solid;
      }
    }
    &__meta-info {
      padding-top: .5rem;
      display: flex;
      flex-direction: column;
      padding-bottom: .5rem;
    }
    &__description {

    }
    &__map {
      height: 320px;
    }
    &__address {
      margin-top: .5rem;
      display: flex;
      &>div {
        width: 50%;
      }
      @media (max-width: 510px) {
        flex-direction: column;
        &>div {
          width: 100%;
        }
      }
    }
    &__location-address {
      display: flex;
    }
  }
  .location-info {
    align-items: center;
    padding: .3rem;
    &__header {
      margin-right: 1rem;
      font-size: 1rem;
      font-weight: bolder;
    }
    &__value {
      color: #5a5c61;
      font-size: 1rem;
    }
    img {
      margin-right: .25rem;
      width: 16px;
    }
  }
  .hidden {
    position: relative;

  }
  .contact {
    &__column {
      display: flex;
      flex-direction: column;
      &--header {
        align-items: flex-end;
        padding-right: .3rem;
        font-weight: bold;
      }
      &--value {
        color: #6c757d;
      }
    }
    &__item {

    }
  }
</style>