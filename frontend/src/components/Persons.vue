<template>
  <div class="personclass">
    <h1>{{ msg }}</h1>
    <input id="personName" type="text" value="Darth" v-model="personName"></input>
    <button class="btn btn-primary" v-on:click="getPerson()">Get Star wars person</button>
    <h2>One fetched person</h2>
    <div>{{ person.person.id }} {{ person.person.name }} {{ person.person.surname }}</div>
    <h2>List of all</h2>
    <div v-for="value in persons.personList">
      {{ value.id }} {{ value.name }} {{ value.surname }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'personclass',
  data () {
    return {
      msg: 'Star Wars persons',
      person: {person: {id: '', name: '', surname: ''}},
      persons: ''
    }
  },
  mounted () {
    this.getPersons()
  },
  methods: {
    getPerson () {
      console.log(this.personName)
      this.$http.get('http://localhost:8000/graphql?query=' + encodeURIComponent('{person(name:"' + this.personName + '"){id,name,surname}}')).then(response => {
        this.person = JSON.parse(response.body).data
      }, response => {
        console.log(response)
      })
    },
    getPersons () {
      this.$http.get('http://localhost:8000/graphql?query=' + encodeURIComponent('{personList{id,name,surname}}')).then(response => {
        this.persons = JSON.parse(response.body).data
      }, response => {
        console.log(response)
      })
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

.personclass {padding-left:35%;}
div {text-align:left;}
</style>
