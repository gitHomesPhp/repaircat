<template>
  <div>
    <div class="sc card">
      <div class="sc__header">
        <div class="sc__title">
          <div class="sc__type">
            <span>{{ title }}</span>
          </div>
          <h1 class="sc__name">{{scCardExtension.sc.name}}</h1>
        </div>
        <div class="sc__contacts contacts">
          <div class="contact__column contact__column--header">
            <span>Телефон:</span>
            <span>Почта:</span>
            <span>Сайт:</span>
          </div>
          <div class="contact__column contact__column--value">
            <span class="contact__item">{{scCardExtension.sc.phone || 'Нет информации'}}</span>
            <span class="contact__item">{{scCardExtension.sc.email || 'Нет информации'}}</span>
            <span class="contact__item">{{scCardExtension.sc.site || 'Нет информации'}}</span>
          </div>
        </div>
      </div>
      <div class="sc__body">
        <div class="sc__meta-info">
          <CommonRating
              :rating="scCardExtension.review_info.rating"
              :count="scCardExtension.review_info.count"
          />
        </div>
        <div class="sc__add-review">
          <button @click="showReviews">Оставить отзыв</button>
        </div>
      </div>
      <div class="sc__address">
        <div>
          <div class="sc__location-address location-info">
            <img src="/img/location_icon.png" alt="">
            <span class="location-info__header">Аддрес:</span>
            <span class="location-info__value">{{ scCardExtension.sc.location.address }}</span>
          </div>
          <div v-if="scCardExtension.sc.location.undergrounds.length"
               v-for="underground in scCardExtension.sc.location.undergrounds"
               class="sc__location-address location-info"
          >
            <img src="/img/underground_icon.png" alt="">
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
    <div class="card reviews" id="reviews">
      <h3 class="reviews__header">Отзывы о &#171{{scCardExtension.sc.name}}&#187</h3>
      <div class="reviews__elements">
        <CommonReview
            v-if="reviewStore.reviews.length"
            v-for="(review, index) in reviewStore.reviews"
            v-show="index < countShowingReviews"
            :key="index"
            :text="review.review.text"
            :rating="review.review.rating"
            :visitor="review.visitor"
        />
        <button @click="showReviews">Показать отзывы</button>
      </div>
    </div>
    <div class="card about" id="information">
      <h3 class="about__header">Информация о сервисном центре</h3>
      <div class="about__description" v-html="scCardExtension.sc.description"></div>
    </div>
  </div>
</template>

<script setup lang="ts">

import {useReviewStore} from "~/stores/reviewStore";

const reviewStore = useReviewStore()
const route = useRoute()

const title = ref('Сервисный центр')
const { data: scCardExtension} = await useFetch(`/api/sc/${route.params.id}`)
const showMap  = ref(false)

const coords = ref([scCardExtension.value.sc.location.latitude, scCardExtension.value.sc.location.longitude])
const changeCoords = (e) => {
  coords.value =  e.get('coords');
}

const countShowingReviews = ref(10)

const showReviews = () => {
  countShowingReviews.value += 10
}
onMounted(() => {
  setTimeout(() => {showMap.value = true}, 150)

})
</script>

<style scoped lang="scss">
  .sc {
    display: flex;
    flex-direction: column;
    margin-bottom: 1rem;
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
      align-items: center;
      @media (max-width: 510px) {
        flex-direction: column-reverse;
      }
    }
    &__add-review {
      @media (max-width: 510px) {
        padding-top: .4rem;
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
  .reviews {
    margin-bottom: 1rem;
    &__header {
      margin-bottom: 1rem;
      font-size: 1.1rem;
      color: #0070c0;
    }
    &__elements {
      font-size: 1rem;
      display: flex;
      flex-direction: column;
    }
  }
  .about {
    display: flex;
    flex-direction: column;
    color: #5a5c61;
    font-family: Roboto,Arial,sans-serif;
    font-size: 1rem;;
    padding-top: .5rem;
    padding-bottom: .5rem;
    &__header {
      margin-bottom: 1rem;
      font-size: 1.1rem;
      color: #0070c0;
    }
    &__description {
      font-size: 1rem;
    }
  }
</style>