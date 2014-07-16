// Activates the Carousel
$('.carousel').carousel({
    interval: 5000
})

// Activates Tooltips for Social Links
$('.tooltip-social').tooltip({
    selector: "a[data-toggle=tooltip]"
})

// Activate Bootstrap Maxlength
$('textarea.form-control').maxlength({
    alwaysShow: true,
    warningClass: "label label-success",
    limitReachedClass: "label label-danger"
});
$('input.form-control').maxlength({
    alwaysShow: true,
    warningClass: "label label-success",
    limitReachedClass: "label label-danger"
});