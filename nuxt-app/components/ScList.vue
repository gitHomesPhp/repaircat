<template>
  <div>
    <ScCard
      v-for="sc in scList"
      :key="sc.id"
      :name="sc.name"
      :id="sc.id"
      :description="sc.description"
      :phone="sc.phone"
      :email="sc.email"
      :site="sc.site"
      :location="sc.location"
      @click="openSc(sc.id)"
    />
  </div>
</template>

<script lang="ts" setup>
import {useScPaginatorStore} from "~/stores/scPaginatorStore";
import {storeToRefs} from "pinia";

const router = useRouter()

const openSc = (id) => {
  //TODO
  router.push(`/spb/sc/${id}-service-center`)
}

const {scList} = storeToRefs(useScPaginatorStore())
</script>

<style scoped>

</style>