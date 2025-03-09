from sklearn.linear_model import LogisticRegression
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score

# Create and test a logistic model
def create_logistic_model(x, y):
    try:
        # Split data into training and testing
        x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.2, random_state=69)

        # Create logistic model and fit it
        model = LogisticRegression(max_iter=250, random_state=69)
        model.fit(x_train, y_train)

        # Predict
        y_pred = model.predict(x_test)

        # Check accuracy
        accuracy = accuracy_score(y_test, y_pred)

        return round(accuracy, 4), None
    except Exception as e:
        return None, str(e)