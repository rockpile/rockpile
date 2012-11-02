(defun foo ()
  1(let ((dummy '(1 2)))
      (mapc #'(lambda (&rest x) (print (append x dummy)))
	    '(1 2 3 5 8)
	    '(3 4 7)
	    '(4 7 8 8 9))))

