$(document).ready(function () {

    // toggle mobile menu
    $('[data-toggle="toggle-nav"]').on('click', function () {
        $(this).closest('nav').find($(this).attr('data-target')).toggleClass('hidden');
        return false;
    });

    // feather icons
    feather.replace();

    // smooth scroll
    var scroll = new SmoothScroll('a[href*="#"]');

    // tiny slider
    $('#slider-1').slick({
        infinite: true,
        prevArrow: $('.prev'),
        nextArrow: $('.next'),
    });

    $('#slider-2').slick({
        dots: true,
        arrows: false,
        infinite: true,
        slidesToShow: 3,
        slidesToScroll: 1,
        autoplay: true,
        autoplaySpeed: 2000,
        centerMode: true,
        customPaging: function (slider, i) {
            return '<div class="bg-white br-round w-1 h-1 opacity-50 mt-5" id=' + i + '> </div>'
        },
        responsive: [{
            breakpoint: 768,
            settings: {
                slidesToShow: 1
            }
        }, ]
    });

    $("#send-form").submit(function(e) {
 
        e.preventDefault();
        var dataString = $(this).serialize();

        $.ajax({
          type: "GET",
          url: "api/message",
          data: dataString,
          success: function () {
            $("#send-div").html("<div id='message'></div>");
            $("#message")
              .html("<h2 class='white fs-m2 fw-700'><i>Message Sent!</i></h2><i data-feather='check' class='w-l5'></i>")
              .hide()
              .fadeIn(1500);
          }
        });
     
        return false;
    });
      
});