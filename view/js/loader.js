

$(document).ready(function() {
    $(window).scroll(function() { //detect page scroll
        var wd = $(window);
        if(wd.scrollTop() + wd.height() >= $(document).height()) { //if user scrolled to bottom of the page
            load();
        }
    });
});
