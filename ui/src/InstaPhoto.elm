module InstaPhoto exposing (main)
import Html exposing (..)
import Html.Attributes exposing (class)
import Html.Events exposing (onClick)
import Browser
import Dict exposing (update)


viewDetailedPhoto : { liked: Bool } -> Html Msg
viewDetailedPhoto  model =
    let 
        buttonType =
            if model.liked then
                "star"
            else
                "star_border"
        msg =
            if model.liked then
                Unlike
            else
                Like
    in
        div [ class "detailed-photo" ]
            [
            div [ class "photo-info" ]
                [ div [ class "like-button" ]
                    [ span 
                        [ class "material-icons md-48"
                        , onClick msg
                        ]
                        [ text buttonType ]
                    ]
                    
                ]
            ]

-- Model

initialModel : { liked : Bool }
initialModel =
    {
        
        liked = False
    }

type Msg = Like | Unlike

update : Msg -> {  liked : Bool } -> { liked : Bool}
update msg model =
    case msg of
        Like ->
            { model | liked = True }
        Unlike ->
            { model | liked = False }
view : {  liked : Bool} -> Html Msg
view model =
    div []
    [ div [ class "header" ]
        [ h1[] [ text "InstaPhoto" ] ]
    ,div [ class "content-flow" ]
        [ viewDetailedPhoto model]
    
    ]
-- Main
main : Program () {  liked : Bool} Msg
main =
    Browser.sandbox{
        init = initialModel
        ,view = view
        ,update = update
    }