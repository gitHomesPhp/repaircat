<template>
  <div class="menu">
    <div class="menu__item">
      <div @click="() => {showScSubMenu = !showScSubMenu}" class="multi-menu">
        <span class="menu__text pointer">
          Сервисный центр
        </span>
        <CommonSubMenu :show="showScSubMenu">
          <ul>
            <li class="menu__item">
              <NuxtLink to="/admin/sc/create" class="menu__text">Добавить</NuxtLink>
            </li>
          </ul>
        </CommonSubMenu>
      </div>
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

<script lang="ts" setup>
  const showScSubMenu = ref(false)
</script>

<style scoped lang="scss">
  .menu {
    display: flex;
    flex-direction: row;
    padding: .4rem .2rem;
    &__item {
      padding: .3rem .2rem;
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
  .multi-menu {
    position: relative;
  }
</style>