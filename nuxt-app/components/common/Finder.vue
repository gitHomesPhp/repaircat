<template>
  <div class="finder">
    <div class="finder__main-search">
      <input @focus="showSub"
             @blur="closeSub"
             type="text"
             placeholder="Метро или район">
      <div v-if="activeSub" class="finder__helper-output">
        <!-- TODO тут отображение -->
      </div>
    </div>

    <button>НАЙТИ</button>
  </div>
</template>

<script lang="ts" setup>
  const activeSub = ref(false)

  const showSub = () => activeSub.value = true
  const closeSub = () => activeSub.value = false
</script>

<style scoped lang="scss">
.finder {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  margin-bottom: .4rem;
  &__main-search {
    position: relative;
    width: 100%;
    input {
      box-sizing: border-box;
      padding: .7rem;
      width: 100%;
      border-radius: 4px;
      border: 1px solid #757575;
      &:focus {
        border-color: #0070c0;
        border-width: 2px;
      }
    }
  }
  &__helper-output {
    position: absolute;
    width: 99.7%;
    border: #5a5c61 1px solid;
    top: 50px;
  }

  button {
    width: 25%;
    height: 90%;
    margin-left: 1rem;
    font-size: 1rem;
    font-weight: bold;
    color: white;
    background-color: #ec2327;
    border: none;
    border-radius: 4px;
    letter-spacing: 1px;
    &:hover {
      background-color: #ce2326;
    }
    &:active {
      background-color: #ab2325;
    }
  }
}
</style>