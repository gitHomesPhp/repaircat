<template>
  <div class="menu">
    <div class="menu__item">
      <NuxtLink to="/admin" style="text-decoration: none;">
        Добавить сервисный центра
      </NuxtLink>
    </div>
    <div class="menu__item">
      <NuxtLink to="/admin/city" style="text-decoration: none;">
        Добавить город
      </NuxtLink>
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
    }
  }
</style>