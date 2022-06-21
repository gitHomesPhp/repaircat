<template>
  <div class="wrapper">
    <img src="~/assets/img/gear_icon.png" alt="">
    <div class="sc-card">
      <h3 class="sc-card__name">
        {{ name }}
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
            <a :href="site">{{ site }}</a>
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
    min-width: 510px;
    margin-top: .4rem;
    margin-bottom: .4rem;
    display: flex;
    flex-direction: row;
    align-items: center;
    &>img {
      width: 100px;
      height: 75px;
    }
  }
  .sc-card {
    cursor: pointer;
    border: 1px solid #eaeaea;
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
    }
    &__name {
      font-weight: bolder;
    }
    &__description {
      padding: .5rem 0;
      font-size: medium;
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