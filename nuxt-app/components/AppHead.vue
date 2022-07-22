<template>
  <div class="header">
    <div class="header__logo">
      <NuxtLink to="/" style="text-decoration: none; color: white">
        <span>Repaircat</span>
      </NuxtLink>
      <div class="header__city">
        {{currentCity.label}}
      </div>
    </div>
    <header>
      <nav class="header__navigation">
        <slot />
      </nav>
    </header>
  </div>
</template>

<script lang="ts" setup>
import {useCityStore} from "~/stores/cityStore";
import {storeToRefs} from "pinia";

const { currentCity } = storeToRefs(useCityStore())
</script>

<style scoped lang="scss">
  .header {
    display: flex;
    flex-direction: column;
    position: sticky;
    top: 0;
    z-index: 999;
    &__logo {
      width: 100%;
      background: #000;
      color: #fff;
      padding: .4rem;
      font-weight: bolder;
      font-size: 1.3rem;
      display: flex;
      justify-content: space-between;
    }
    &__city {
      margin-right: .8em;
    }
    &__navigation {
      background-color: #f2f8fc;
    }
  }
</style>