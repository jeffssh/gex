<template>
  <div v-if="isActive" v-show="isActive" div class="Tab">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'Tab',
  props: {
    name: { required: true },
    selected: { default: false },
    data: { default: '' }
  },
  data() {
    return { isActive: false }
  },
  computed: {
    href() {
      return '#' + this.name.toLowerCase().replace(/ /g , '-');
    }
  },
  mounted() {
    this.isActive = this.selected;
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
