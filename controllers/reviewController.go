func AddAReview() gin.HandlerFunc {
        return func(c *gin.Context) {
                if err := helper.CheckUserType(c, "USER"); err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                        return
                }
                ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
                var review models.Reviews
                defer cancel()

                //validate the request body
                if err := c.BindJSON(&review); err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{
                                "Status":  http.StatusBadRequest,
                                "Message": "error",
                                "Data":    map[string]interface{}{"data": err.Error()}})
                        return
                }
                //use the validator library to validate required fields
                if validationErr := validate.Struct(&review); validationErr != nil {
                        c.JSON(http.StatusBadRequest, gin.H{
                                "Status":  http.StatusBadRequest,
                                "Message": "error",
                                "Data":    map[string]interface{}{"data": validationErr.Error()}})
                        return
                }

                newReview := models.Reviews{
                        Id:          primitive.NewObjectID(),
                        Movie_id:    review.Movie_id,
                        Reviewer_id: review.Reviewer_id,
                        Review:      review.Review,
			Review_id: review.Review_id,
                }

                result, err := reviewCollection.InsertOne(ctx, newReview)

                if err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{
                                "Status":  http.StatusBadRequest,
                                "Message": "error",
                                "Data":    map[string]interface{}{"data": err.Error()}})
                        return
                }

                if err != nil {
                        c.JSON(http.StatusInternalServerError, gin.H{
                                "Status":  http.StatusInternalServerError,
                                "Message": "error",
                                "Data":    map[string]interface{}{"data": err.Error()}})
                        return
                }

                c.JSON(http.StatusCreated, gin.H{
                        "Status":  http.StatusCreated,
                        "Message": "success",
                        "Data":    map[string]interface{}{"data": result}})
        }
}
