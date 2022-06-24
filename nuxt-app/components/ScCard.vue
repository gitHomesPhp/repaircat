<template>
  <div class="wrapper">
    <img src="~/assets/img/gear_icon.png" alt="">
    <div class="sc-card">
      <h3 class="sc-card__name">
        <!--TODO-->
        <NuxtLink @click.prevent.stop :to="`/spb/sc/${id}-service-center`" style="text-decoration: none; color: #0070c0">{{ name }}</NuxtLink>
      </h3>
      <div class="sc-card__description" v-html="description"></div>
      <div class="sc-card__bottom">
        <div class="sc-card__location location">
          <div class="location__item">
            <span class="location__image">
              <img src="~/assets/img/location_icon.png" alt="">
            </span>
            <span>
              {{ location.address }}
            </span>
          </div>
          <div class="location__item">
            <span class="location__image">
              <img src="~/assets/img/underground_icon.png" alt="">
            </span>
            <span>
              {{ location.underground.label }}
            </span>
          </div>
        </div>
        <div class="sc-card__contact contact">
          <div class="sc-card__email contact__item">
            <span class="contact__image">
              <img src="~/assets/img/email_icon.png" alt="">
            </span>
            <span>
              {{ email }}
            </span>
          </div>
          <div class="sc-card__site contact__item">
          <span class="contact__image">
            <img src="~/assets/img/site_icon.png" alt="">
          </span>
          <span>
            <a @click.stop :href="site" target="_blank">{{ site }}</a>
          </span>
          </div>
          <div class="sc-card__phone contact__item">
          <span class="contact__image">
            <img src="~/assets/img/phone_icon.png" alt="">
          </span>
            <span>
            {{ phone }}
          </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  defineProps({
    id: { required: true, type: Number },
    name: { required: true, type: String },
    description: { required: true, type: String },
    phone: { required: true, type: String },
    email: { required: true, type: String },
    site: { required: true, type: String },
    city: { default: '', type: String },
    location: { required: true, type: Object },
  })
</script>

<style scoped lang="scss">
  .wrapper {
    margin-top: .4rem;
    margin-bottom: .4rem;
    display: flex;
    flex-direction: row;
    align-items: center;
    @media(max-width: 510px) {
      flex-direction: column;
      &>img {
        margin-bottom: .4rem;
      }
    }
    &>img {
      width: 100px;
      height: 75px;
    }
  }
  .sc-card {
    cursor: pointer;
    border: 1px solid #eaeaea;
    background-color: white;
    padding: 15px;
    border-radius: 5px;
    text-align: left;
    flex: 1 1;
    display: flex;
    flex-direction: column;
    transition: box-shadow .2s ease,border .2s ease;
    &:hover {
      border-radius: var(--geist-radius);
      text-align: left;
      flex: 1 1;
      display: flex;
      flex-direction: column;
      transition: box-shadow .2s ease;
      box-shadow: 0 8px 30px rgba(0,0,0,.12);
      border: 1px solid transparent;
    }
    &__bottom {
      display: flex;
      justify-content: space-between;
      font-size: .9rem;
      @media(max-width: 510px) {
        flex-direction: column;
      }
    }
    &__name {
      font-weight: bolder;
      font-size: 1.5rem;
    }
    &__description {
      padding: .5rem 0;
      font-size: medium;
      overflow: hidden;
      position: relative;
      height: 112px;
      color: #5a5c61;
      &:after {
        content: "";
        text-align: right;
        position: absolute;
        bottom: 0;
        right: 0;
        left: 0;
        height: 6.5rem;
        background: linear-gradient(to bottom, rgba(255, 255, 255, 0), white 100%);
        pointer-events: none;
      }
    }
    &__contact {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      align-items: flex-end;
    }
    &__email {
      display: flex;
      align-items: center;
    }
    &__site {
      display: flex;
      align-items: center;
      a {
        text-decoration: none;
        color: #0070c0;
      }
    }
    &__phone {
      display: flex;
      align-items: center;
    }
    .contact {
      &__image {
        margin-right: .3rem;
        img {
          width: 20px;
        }
      }
      &__item {
        margin-bottom: .4rem;
      }
    }
    .location {
      &__image {
        margin-right: .3rem;
        img {
          width: 20px;
        }
      }
      &__item {
        margin-bottom: .4rem;
      }
    }
  }
</style>