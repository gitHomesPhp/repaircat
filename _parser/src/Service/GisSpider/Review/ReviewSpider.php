<?php

declare(strict_types=1);

namespace App\Service\GisSpider\Review;

use App\Entity\Sc;
use App\Entity\Url;
use DateTimeImmutable;
use Doctrine\ORM\EntityManagerInterface;
use GuzzleHttp\Client;
use stdClass;

class ReviewSpider
{
    private Review $review;

    public function __construct(
        private Client $client,
        private EntityManagerInterface $em,
    ) {}

    public function handle(?int $projectScId = null)
    {
        if ($projectScId) {
            $scRepo = $this->em->getRepository(Sc::class);
            $sc = $scRepo->findOneBy(['repaircat_id' => $projectScId]);

            //TODO не факт что в базе в будущем будет инт
            $_2gisId = (int)$sc->getGisId();

            $link = new Link();
            $this->dig($link->setItemId($_2gisId)->buildUrl());

        }

    }

    //public function setReview(Review $review): self
    //{
    //    $this->review = $review;
//
    //    return $this;
    //}

    public function dig(Link $link): void
    {
        $response = $this->client->get($link->getUrl());

        $url = new Url();
        $url->setHost(Link::HOST);
        $url->setType(Url::REVIEW_TYPE);
        $url->setPath($link->getPath());
        $url->setLastRequest(new DateTimeImmutable());

        if ($response->getStatusCode() === 200) {
            $data = $response->getBody()->getContents();
            $url->setIsResponseOK(true);
            $url->setData($data);
        }

        $this->em->persist($url);
        $this->em->flush();
    }

    private function getScReviewNextLink(stdClass $dataObject)
    {

    }
}