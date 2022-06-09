<template>
  <div class="header">
    <div class="header__logo">
      <span>Repaircat</span>
    </div>
    <header>

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
  }
</style>