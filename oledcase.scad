// https://openscad.org/
module screen(){
    union() {
        cube([28,28,6]);
        translate([9,0,6]) cube([10,2,6]);
    }
    
}
module arduino_pro(){
    union() {
        cube([18,33,4]);
        translate([5.5,-2,1.5]) cube([7,6,3]);
    }
}
//screen();
//translate([30,0,0]) arduino_pro();

module screen_and_arduino(){
    translate([5.5,0,0]){
        translate([28-5.5,42,28]) rotate([90,180,0]) screen();
        arduino_pro();
    }
}
module case(){
    difference(){
        translate([0,2,0])
        difference() {
            cube([34,42,34]);
            translate([3,3,3]) cube([29,28,29]);
        }
        translate([3,3,3]) !screen_and_arduino();
    }
}
module cased(){
    difference(){
        cube([37,37,44]);
        translate([4,4,4]) cube([29,29,42]);
        #translate([4.5,4.5,1.5]) rotate([90,0,90]) screen_and_arduino();
        translate([20,28,1]) rotate([180, 0, 90])
            linear_extrude(1)
                text("TIO", halign="right", size=8);
    }
}
cased();