var currentTab;
var composeCount = 0;
//initilize tabs
$(function() {
    //when any tab is clicked this method will fire
    $("#myTab").on("click", "a", function (e) {
        e.preventDefault();

        $(this).tab('show');
        $currentTab = $(this);
    });


    

    registerOpenButtonEvent();
    registerOpenProjectEvent();
    registerCloseEvent();
});

//this method will demonstrate how to add tab dynamically
function registerOpenButtonEvent() {
    /* just for this demo */
    $('#openButton').click(function (e) {

        console.log("clicked open")
        e.preventDefault();
        $('#projectDir').click()
        
    });

}

function registerOpenProjectEvent() {
    $("input:file").change(function (){
        var fileName = $("input:file")[0].files[0].name;
        
        var projectPath = $("input:file")[0].files[0].path        

        var tabName = "" + fileName; 

        let tabIDGenerator = function () {
            let S4 = function() {
                return (((Date.now()))|0).toString(16).substring(1);
            };
            return (S4()+S4()+"-"+S4()+"-"+S4()+"-"+S4()+"-"+S4()+S4()+S4());
        };

        var tabId = tabIDGenerator()

        $('.nav-tabs').append('<li><a href="#' + tabId + '"><button class="close closeTab" type="button" >×</button>' + tabName + '</a></li>');
        $('.tab-content').append('<div class="tab-pane" id="' + tabId + '"></div>');

        craeteNewTabAndLoadUrl("", projectPath, "#" + tabId);

        $(this).tab('show');
        showTab(tabId);
        registerCloseEvent();
    });
}

//this method will register event on close icon on the tab..
function registerCloseEvent() {

    $(".closeTab").click(function () {

        //there are multiple elements which has .closeTab icon so close the tab whose close icon is clicked
        var tabContentId = $(this).parent().attr("href");
        $(this).parent().parent().remove(); //remove li of tab
        $('#myTab a:last').tab('show'); // Select first tab
        $(tabContentId).remove(); //remove respective tab content

    });
}

//shows the tab with passed content div id..paramter tabid indicates the div where the content resides
function showTab(tabId) {
    $('#myTab a[href="#' + tabId + '"]').tab('show');
}
//return current active tab
function getCurrentTab() {
    return currentTab;
}

//This function will create a new tab here and it will load the url content in tab content div.
function craeteNewTabAndLoadUrl(parms, projectPath, loadDivSelector) {

    // $("" + loadDivSelector).load(url, function (response, status, xhr) {
            $(loadDivSelector).html(""+projectPath);

        // if (status == "error") {
        //     var msg = "Sorry but there was an error getting details ! ";
        //     $("" + loadDivSelector).html(msg + xhr.status + " " + xhr.statusText);
        // }
    // });

}

//this will return element from current tab
//example : if there are two tabs having  textarea with same id or same class name then when $("#someId") whill return both the text area from both tabs
//to take care this situation we need get the element from current tab.
function getElement(selector) {
    var tabContentId = $currentTab.attr("href");
    return $("" + tabContentId).find("" + selector);

}


function removeCurrentTab() {
    var tabContentId = $currentTab.attr("href");
    $currentTab.parent().remove(); //remove li of tab
    $('#myTab a:last').tab('show'); // Select first tab
    $(tabContentId).remove(); //remove respective tab content
}