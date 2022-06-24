<template>
  <div class="menu">
    <div class="menu__item">
      <NuxtLink to="/admin" class="menu__text">
        Сервисный центр
      </NuxtLink>
    </div>
    <div class="menu__item">
      <NuxtLink to="/admin/city" class="menu__text">
        Словари
      </NuxtLink>
    </div>
    <div class="menu__item">
      <span class="menu__text pointer">
        Тестовый пункт
      </span>
    </div>
  </div>
</template>

<script lang="ts">

</script>

<style scoped lang="scss">
  .menu {
    display: flex;
    flex-direction: row;
    padding: .4rem .2rem;
    &__item {
      padding: 0 .2rem;
      font-size: .9rem;
    }
    &__text {
      text-decoration: none;
      color: cadetblue;
      &:hover {
        color: #367a7c;
      }
    }
  }
</style>