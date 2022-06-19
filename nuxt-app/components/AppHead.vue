<template>
  <div class="header">
    <div class="header__logo">
      <NuxtLink to="/" style="text-decoration: none; color: white">
        <span>Repaircat</span>
      </NuxtLink>
    </div>
    <header>
      <nav class="header__navigation">
        <slot />
      </nav>
    </header>
  </div>
</template>

<script>
export default {
  name: "Head"
}
</script>

<style scoped lang="scss">
  .header {
    display: flex;
    flex-direction: column;
    margin-bottom: 1rem;
    &__logo {
      width: 100%;
      background: #000;
      color: #fff;
      padding: .4rem;
      font-weight: bolder;
      font-size: 1.3rem;
    }
    &__navigation {
      background-color: #f2f8fc;
    }
  }
</style>