(defclass people () ((name :initarg :name) (age :initarg :age :initform 0)))
(defclass rockpile (people) ())

(defgeneric get_name (people) (:documentation "get people name"))

(defparameter *badguy* (make-instance 'people :name "badguy" :age 30))