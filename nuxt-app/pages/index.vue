<template>
  <div class="container">
    <AppHead />
    <AppBody>
      <ScList />
    </AppBody>
    <AppFoot />
  </div>
</template>

<script>
import ScList from "../components/ScList";
export default {
  name: "index",
  components: {ScList}
}
</script>

<style scoped lang="scss">

</style>