<template>
  <form @submit.prevent="goToSearchScList" class="finder">
    <div class="finder__main-search">
      <input @focus="showSub"
             @blur="closeSub"
             type="text"
             placeholder="Метро или район"
             v-model="findString"
      >
      <div v-if="activeSub" class="finder__helper-output">
        <div class="empty" v-if="submenuItems.empty
                  && submenuItems.undergrounds.length === 0
                  && submenuItems.municipalities.length === 0"
        >
          <span>{{submenuItems.empty.label}}</span>
        </div>
        <ul>
          <li v-if="submenuItems.undergrounds.length"
              v-for="underground in submenuItems.undergrounds"
              class="find-by"
              @click="goToUndergroundsScList(underground)"
          >
            <img src="/img/underground.svg" alt="">
            {{underground.label}}
          </li>
          <li v-if="submenuItems.municipalities.length"
              v-for="municipality in submenuItems.municipalities"
              class="find-by"
              @click="goToMunicipalityScList(municipality)"
          >
            <img src="/img/municipality.png" alt="">
            {{municipality.label}}
          </li>
        </ul>
      </div>
    </div>

    <button>НАЙТИ</button>
  </form>
</template>

<script lang="ts" setup>
  import { useCityStore } from "~/stores/cityStore";
  const { findUndergrounds, findMunicipalities, currentCity } = useCityStore()


  const router = useRouter()
  const emptySubmenuItem = {
    code: 'empty',
    label: 'Вы можете искать по метро или району...'
  }
  const findString = ref('')
  const activeSub = ref(false)
  const foundUndergrounds = ref([])
  const foundMunicipalities = ref([])
  const submenuItems = ref({
    empty: emptySubmenuItem,
    undergrounds: [],
    municipalities: [],
  })

  const showSub = () => activeSub.value = true
  const closeSub = () => setTimeout(() => {activeSub.value = false}, 200)

  const isGoing = ref(false);

  const changeHelperFind = async () => {
    if (!isGoing.value) {
      isGoing.value = true
      setTimeout(() => {
        foundUndergrounds.value = findUndergrounds(findString.value.toLowerCase(), currentCity.id).slice(0, 3)
        submenuItems.value.undergrounds = foundUndergrounds.value
        foundMunicipalities.value = findMunicipalities(findString.value.toLowerCase(), currentCity.id).slice(0, 3)
        submenuItems.value.municipalities = foundMunicipalities.value
        isGoing.value = false
      }, 100)
    }
  }

  const goToUndergroundsScList = async (underground) => {
    await router.push(`/${currentCity.code}/service-center/underground-${underground.slug}`)
  }

  const goToMunicipalityScList = async (municipality) => {
    await router.push(`/${currentCity.code}/service-center/municipality-${municipality.slug}`)
  }

  const goToSearchScList = async () => {
    await router.push(`/${currentCity.code}/service-center?q=${findString.value}`)
  }

  watch(() => findString.value, async () => {
    await changeHelperFind()
  })
</script>

<style scoped lang="scss">
.empty {
  font-size: .9rem;
}
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
    top: 50px;
    z-index: 2;
    background: #fff;
    box-shadow: 0 6px 10px rgb(0 0 0 / 7%), 0 1px 18px rgb(0 0 0 / 6%), 0 3px 5px rgb(0 0 0 / 10%);
    border-radius: 5px;
    padding: .8rem 0;
    font-size: 16px;
    &>div {
      display: flex;
      justify-content: center;
      width: 100%;
    }
    &>ul {
      display: flex;
      flex-direction: column;
      justify-content: center;
    }
  }
  .find-by {
    cursor: pointer;
    display: flex;
    align-items: center;
    font-size: .9rem;
    padding: .4rem .7rem;
    & > img {
      width: 25px;
      margin-right: .7rem;
    }
    &:hover {
      background-color: #f7f7f7;
    }
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