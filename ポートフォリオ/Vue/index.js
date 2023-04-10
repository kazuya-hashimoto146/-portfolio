var app = new Vue({
    el: '#app',
    data: {
        message: 'Hello vue!'
    },
    methods: {
        changeMessage: function() {
            if (this.message === 'Hello vue!'){
                this.message = 'Welcome!'    
            } else {
                this.message = 'Hello vue!'
            }
        
        }
    }
})