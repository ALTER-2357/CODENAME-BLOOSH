
    function cus(){  
      window.location.href = "./contact-us.html";
        }  








        function formSubmit(event) {
          var url = "http://localhost:3030/upload";
          var request = new XMLHttpRequest();
          request.open('POST', url, true);
          request.onload = function() { // request successful
          // we can use server response to our request now
            console.log(request.responseText);
          };
          request.onerror = function() {
            // request failed
          };
          request.send(new FormData(event.target)); // create FormData from form that triggered event
          event.preventDefault();
        }
        // and you can attach form submit event like this for example
        function attachFormSubmitEvent(formId){
          document.getElementById(formId).addEventListener("submit", formSubmit);
        } 
        
        
        function jobs(){ 
            window.location.href = "./jobs.html";
        
        }
        
            function openNav() {
              document.getElementById("mySidenav").style.width = "100%";
            }
            
            function closeNav() {
              document.getElementById("mySidenav").style.width = "0";
            }
          
          
