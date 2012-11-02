(defun sbcl-socket-connect (port)
  (let ((socket (make-instance 'sb-bsd-sockets:inet-socket :type :stream :protocol :tcp)))

    (sb-bsd-sockets:socket-connect 
     socket 
     (sb-bsd-sockets:host-ent-address (sb-bsd-sockets:get-host-by-name "localhost")) 
     port)

    (sb-bsd-sockets:socket-make-stream 
     socket 
     :input t 
     :output t 
     :buffering 
     :line 
     :element-type 'character)
    
    (sb-bsd-sockets:socket-send socket (concatenate 'string "123" (list #\return #\newline)) 5)
    (print "this")
    (print (with-input-from-string (istream (sb-bsd-sockets:socket-receive socket nil 65535)) (read istream nil nil)))
    (sb-bsd-sockets:socket-close socket)))

