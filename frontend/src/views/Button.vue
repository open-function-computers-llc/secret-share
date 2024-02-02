<template>
    <div v-if="!errorMessage">
       <a class = "button" :href = "'/#/show=true/' + secretID">Click here to see your secret</a>
    </div>

    <div v-else>
        <h2>Sorry, that secret could not be found. 😢</h2>
    </div>
</template>

<script>
export default {
    data() {
        return {
            secretID: this.$route.params.id,
            errorMessage: "",
        }
    },
}
</script>
<style lang="css" scoped>
.button {
  color: white;
  text-decoration: none;
  background-color: #6f4bf3;
  padding: 5px 15px;
  display: block;
  transition: all 0.3s ease, filter 0s linear;
  border: 0;
  font-size: 1.25rem;
  border-radius: 3px;
  cursor: pointer;
  margin: auto;
  width: 30%;
}
li.active .button {
  background-color: white;
  color: #6f4bf3;
}

.deactived {
  opacity: 0.25;
}

.button-wrapper {
  display: flex;
  justify-content: center;
  margin-top: 60px;
}

button.button {
  text-transform: uppercase;
  font-size: 1.5rem;
  padding: 10px 30px;
}

.button:hover {
  filter: brightness(1.1);
}
</style>