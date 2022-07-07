<?php

declare(strict_types=1);

namespace App\Service\GisSpider\Review;

use GuzzleHttp\Client;

class ReviewSpider
{
    public function __construct(
        private Client $client,
        //private ?Review $review = null,
    ) {}

    //public function setReview(Review $review): self
    //{
    //    $this->review = $review;
//
    //    return $this;
    //}
//
    //public function dig(): void
    //{
    //    $response = $this->client->get($this->review->getUrl());
    //}
}