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
          <span>{{sc.phone}}</span>
          <span>{{sc.email}}</span>
          <span>{{sc.site}}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  const route = useRoute()

  const title = ref('Сервисный центр')
  const { data: sc} = await useFetch(`/api/sc/1`)
</script>

<style scoped lang="scss">
  .sc {
    display: flex;
    flex-direction: column;
    &__header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding-bottom: 1rem;
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
  }
</style>